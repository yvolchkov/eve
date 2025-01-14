// Copyright (c) 2017-2018 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Allocate a small integer for each Network UUID.
// Persist the numbers across reboots using uuidtonum package
// When there are no free numbers then reuse the unused numbers.

package zedrouter

import (
	"github.com/lf-edge/eve/pkg/pillar/types"
	"github.com/lf-edge/eve/pkg/pillar/uuidtonum"
	"github.com/satori/go.uuid"
)

var AllocReservedBridgeNumBits types.Bitmap

// Read the existing bridgeNums out of what we published/checkpointed.
// Also read what we have persisted before a reboot
// Store in reserved map since we will be asked to allocate them later.
// Set bit in bitmap.
func bridgeNumAllocatorInit(ctx *zedrouterContext) {

	pubNetworkInstanceStatus := ctx.pubNetworkInstanceStatus
	pubUuidToNum := ctx.pubUuidToNum

	items := pubUuidToNum.GetAll()
	for _, st := range items {
		status := st.(types.UuidToNum)
		if status.NumType != "bridgeNum" {
			continue
		}
		log.Functionf("bridgeNumAllocatorInit found %v\n", status)
		bridgeNum := status.Number
		uuid := status.UUID
		// If we have a config for the UUID we should mark it as
		// allocated; otherwise mark it as reserved.
		// XXX however, on startup we are not likely to have any
		// config yet.
		if AllocReservedBridgeNumBits.IsSet(bridgeNum) {
			log.Errorf("AllocReservedBridgeNums already set for %d\n",
				bridgeNum)
			continue
		}
		log.Functionf("Reserving bridgeNum %d for %s\n", bridgeNum, uuid)
		AllocReservedBridgeNumBits.Set(bridgeNum)
		// Clear InUse
		uuidtonum.UuidToNumFree(log, ctx.pubUuidToNum, uuid)
	}
	// In case zedrouter process restarted we fill in InUse from
	// NetworkInstanceStatus
	items = pubNetworkInstanceStatus.GetAll()
	for _, st := range items {
		status := st.(types.NetworkInstanceStatus)
		bridgeNum := status.BridgeNum
		uuid := status.UUID

		// If we have a config for the UUID we should mark it as
		// allocated; otherwise mark it as reserved.
		// XXX however, on startup we are not likely to have any
		// config yet.
		if !AllocReservedBridgeNumBits.IsSet(bridgeNum) {
			log.Fatalf("AllocReservedBridgeNumBits not set for %s num %d\n",
				uuid.String(), bridgeNum)
			continue
		}
		log.Functionf("Marking InUse bridgeNum %d for %s\n", bridgeNum, uuid)
		// Set InUse
		uuidtonum.UuidToNumAllocate(log, ctx.pubUuidToNum, uuid, bridgeNum,
			false, "bridgeNum")
	}
}

// If an entry is not inUse and and its CreateTime were
// before the agent started, then we free it up.
func bridgeNumAllocatorGC(ctx *zedrouterContext) {

	pubUuidToNum := ctx.pubUuidToNum

	log.Functionf("bridgeNumAllocatorGC")
	freedCount := 0
	items := pubUuidToNum.GetAll()
	for _, st := range items {
		status := st.(types.UuidToNum)
		if status.NumType != "bridgeNum" {
			continue
		}
		if status.InUse {
			continue
		}
		if status.CreateTime.After(ctx.agentStartTime) {
			continue
		}
		log.Functionf("bridgeNumAllocatorGC: freeing %+v", status)
		bridgeNumFree(ctx, status.UUID, false)
		freedCount++
	}
	log.Functionf("bridgeNumAllocatorGC freed %d", freedCount)
}

func bridgeNumAllocate(ctx *zedrouterContext, uuid uuid.UUID) int {

	// Do we already have a number?
	bridgeNum, err := uuidtonum.UuidToNumGet(log, ctx.pubUuidToNum, uuid,
		"bridgeNum")
	if err == nil {
		log.Functionf("Found allocated bridgeNum %d for %s\n",
			bridgeNum, uuid)
		if !AllocReservedBridgeNumBits.IsSet(bridgeNum) {
			log.Fatalf("AllocReservedBridgeNumBits not set for %d\n",
				bridgeNum)
		}
		// Set InUse and update time
		uuidtonum.UuidToNumAllocate(log, ctx.pubUuidToNum, uuid, bridgeNum,
			false, "bridgeNum")
		return bridgeNum
	}

	// Find a free number in bitmap
	// XXX could look for non-0xFF bytes first for efficiency
	bridgeNum = 0
	for i := 1; i < 256; i++ {
		if !AllocReservedBridgeNumBits.IsSet(i) {
			bridgeNum = i
			log.Functionf("Allocating bridgeNum %d for %s\n",
				bridgeNum, uuid)
			break
		}
	}
	if bridgeNum == 0 {
		log.Functionf("Failed to find free bridgeNum for %s. Reusing!\n",
			uuid)
		oldUuid, oldBridgeNum, err := uuidtonum.UuidToNumGetOldestUnused(log, ctx.pubUuidToNum, "bridgeNum")
		if err != nil {
			log.Fatal("All 255 bridgeNums are in use!")
		}
		log.Functionf("Reuse found bridgeNum %d for %s. Reusing!\n",
			oldBridgeNum, oldUuid)
		uuidtonum.UuidToNumDelete(log, ctx.pubUuidToNum, oldUuid)
		AllocReservedBridgeNumBits.Clear(oldBridgeNum)
		bridgeNum = oldBridgeNum
	}
	if AllocReservedBridgeNumBits.IsSet(bridgeNum) {
		log.Fatalf("AllocReservedBridgeNums already set for %d\n",
			bridgeNum)
	}
	AllocReservedBridgeNumBits.Set(bridgeNum)
	uuidtonum.UuidToNumAllocate(log, ctx.pubUuidToNum, uuid, bridgeNum, true,
		"bridgeNum")
	return bridgeNum
}

func bridgeNumFree(ctx *zedrouterContext, uuid uuid.UUID, fatal bool) {

	bridgeNum, err := uuidtonum.UuidToNumGet(log, ctx.pubUuidToNum, uuid, "bridgeNum")
	if err != nil {
		if fatal {
			log.Fatalf("bridgeNumFree: num not found for %s\n",
				uuid.String())
		} else {
			log.Warnf("bridgeNumFree: num not found for %s\n",
				uuid.String())
		}
		return
	}
	// Check that number exists in the allocated numbers
	if !AllocReservedBridgeNumBits.IsSet(bridgeNum) {
		if fatal {
			log.Fatalf("bridgeNumFree: AllocReservedBridgeNumBits not set for %d\n",
				bridgeNum)
		} else {
			log.Warnf("bridgeNumFree: AllocReservedBridgeNumBits not set for %d\n",
				bridgeNum)
		}
	} else {
		AllocReservedBridgeNumBits.Clear(bridgeNum)
	}
	uuidtonum.UuidToNumDelete(log, ctx.pubUuidToNum, uuid)
}
