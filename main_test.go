package main

import (
	"testing"
)

func Test_getBlobPath(t *testing.T) {
	type args struct {
		basePath string
		digest   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal test", args{"/Users/test/Documents/reg_storage", "sha256:c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2"}, "/Users/test/Documents/reg_storage/blobs/sha256/c3/c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2/data"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBlobPath(tt.args.basePath, tt.args.digest); got != tt.want {
				t.Errorf("getBlobPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBasePath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"normal test", args{"/Users/test/Documents/reg_storage/docker/registry/v2/repositories/library/ubuntu/manifests/revisions/sha256/c3ab8ff13720e8ad9047dd39466b3c8974e592c2fa383d4a3960714caef0c4f2/link"}, "/Users/test/Documents/reg_storage/docker/registry/v2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBasePath(tt.args.path); got != tt.want {
				t.Errorf("getBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
