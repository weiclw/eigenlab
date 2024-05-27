package main

import (
    runner "eigenlab/runner"
)

func getCmdsReal() []string {
    cmdlines := []string{
        "/opt/homebrew/bin/qemu-system-aarch64",
        "-L",
        "/Applications/UTM.app/Contents/Resources/qemu",
        "-cpu",
        "host",
        "-smp",
        "cpus=2,sockets=1,cores=2,threads=1",
        "-machine",
        "virt,",
        "-accel",
        "hvf",
        "-accel",
        "tcg,tb-size=512",
        "-drive",
        "if=pflash,format=raw,unit=0,file=/Applications/UTM.app/Contents/Resources/qemu/edk2-aarch64-code.fd,readonly=on",
        "-drive",
        "if=pflash,unit=1,file=/Applications/UTM.app/Contents/Resources/qemu/edk2-arm-vars.fd",
        "-nographic",
        "-boot",
        "menu=on",
        "-m",
        "2048",
        "-device",
        "intel-hda",
        "-device",
        "virtio-blk-pci,drive=drive0,bootindex=0",
        "-drive",
        "if=none,media=disk,id=drive0,file=/tmp/alpine-virt-3.18.4-aarch64.iso",
        "-device",
        "virtio-blk-pci,drive=drive1,bootindex=1",
        "-drive",
        "if=none,media=disk,id=drive1,file=/tmp/1.qcow2"}

    return cmdlines
}

func main() {
    args := getCmdsReal()
    runner.Run(args)
}
