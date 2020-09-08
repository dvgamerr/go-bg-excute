package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println(os.Args[0])

	// waifu2x-ncnn-vulkan_waifu2xEX -i "C:\Users\dvgamerr\Downloads\83794835_p001_.jpg" -o "C:\Users\dvgamerr\Downloads\83794835_p001_jpg.png" -f png
	cmd := exec.Command("C:\\tools\\waifu2x-gui\\waifu2x-extension-gui\\waifu2x-ncnn-vulkan\\waifu2x-ncnn-vulkan_waifu2xEX.exe", "-i", os.Args[0], "-o", os.Args[0], "-f", "png")
	err := cmd.Start()
	fmt.Println(err)

	// cmd.Stdin=c;
	//cmd.Stdout= c;

	//cmd.Stderr= os.Stderr;
	//cmd.Run();
	//fmt.Println(cmd.Output())
	//fmt.Println(cmd.Stdout)

}
