//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// SPDX-License-Identifier: Apache-2.0
// Copyright 2017-2021 Authors of Cilium

// Code generated by main. DO NOT EDIT.

package kafka

// DeepEqual is an autogenerated deepequal function, deeply comparing the
// receiver with other. in must be non-nil.
func (in *PortRule) DeepEqual(other *PortRule) bool {
	if other == nil {
		return false
	}

	if in.Role != other.Role {
		return false
	}
	if in.APIKey != other.APIKey {
		return false
	}
	if in.APIVersion != other.APIVersion {
		return false
	}
	if in.ClientID != other.ClientID {
		return false
	}
	if in.Topic != other.Topic {
		return false
	}

	return true
}
