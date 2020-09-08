package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	if os.Args[0] == "" {
		return
	}
	path := filepath.Dir(os.Args[1])
	file := filepath.Base(os.Args[1])
	output := path + "\\" + strings.TrimSuffix(file, filepath.Ext(file)) + "_waifu2x.png"

	// waifu2x-ncnn-vulkan_waifu2xEX -i "C:\Users\dvgamerr\Downloads\83794835_p001_.jpg" -o "C:\Users\dvgamerr\Downloads\83794835_p001_jpg.png" -f png
	cmd := exec.Command("C:\\tools\\waifu2x-gui\\waifu2x-extension-gui\\waifu2x-ncnn-vulkan\\waifu2x-ncnn-vulkan_waifu2xEX.exe", "-i", os.Args[1], "-o", output, "-f", "png")
	cmd.Start()

	// cmd.Stdin=c;
	//cmd.Stdout= c;

	//cmd.Stderr= os.Stderr;
	//cmd.Run();
	//fmt.Println(cmd.Output())
	//fmt.Println(cmd.Stdout)

}
