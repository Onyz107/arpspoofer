package arp

import (
	"github.com/google/gopacket/layers"
)

type ARPType uint16

const (
	ARPRequest ARPType = layers.ARPRequest
	ARPReply   ARPType = layers.ARPReply
)
