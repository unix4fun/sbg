/*
 *
 * Copyright 2017 (c) unix4fun.net
 *
 * Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following
 * conditions are met:
 *
 * 1. Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 *
 * 2. Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following
 *    disclaimer in the documentation and/or other materials provided with the distribution.
 *
 * 3. Neither the name of the copyright holder nor the names of its contributors may be used to endorse or promote products derived
 *    from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING,
 * BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT
 * SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
 * OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// BUBBLE fingerprint
// ssh-keygen -B  -l  -f /etc/ssh/ssh_host_rsa_key.pub
//
// SHA256 fingerprint
// ssh-keygen -E sha256  -l  -f /etc/ssh/ssh_host_rsa_key.pub
//
// ascii art stuff
// ssh-keygen -v -l  -f /etc/ssh/ssh_host_rsa_key.pub
//
// first line is ALWAYS the hash
// then comes the ascii art..

const (
	sbgVersion     = 2017072800
	KGCMD          = "/usr/bin/ssh-keygen"
	KGARGS_BUBBLE  = "-B"
	KGARGS_V       = "-v"
	KGARGS_L       = "-l"
	KGARGS_F       = "-f"
	KGARGS_KEYFILE = "/etc/ssh/ssh_host_rsa_key.pub"
	KGOFFSET_FP    = 0
	KGFMT_SEP      = "  "
)

// keyfile 2 ascii buffer
func sshkf2b(kgbin, keyfile string, bubble bool) (fpbuf []string, err error) {
	var cmd *exec.Cmd
	if bubble {
		cmd = exec.Command(kgbin, KGARGS_V, KGARGS_BUBBLE, KGARGS_L, KGARGS_F, keyfile)
	} else {
		cmd = exec.Command(kgbin, KGARGS_V, KGARGS_L, KGARGS_F, keyfile)
	}
	fp, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	fpbuf = strings.Split(string(fp), "\n")
	return
}

func printBanner(fpBufArgs [][]string) {
	// index 0 == FP
	// rest == image
	if len(fpBufArgs) >= 1 {
		fmt.Printf("\n")
		for i := 1; i < len(fpBufArgs[0]); i++ {
			for j := 0; j < len(fpBufArgs); j++ {
				fmt.Printf("%s%s", fpBufArgs[j][i], KGFMT_SEP)
			}
			fmt.Printf("\n")
		}

		for _, v := range fpBufArgs {
			fmt.Printf("%s\n", v[KGOFFSET_FP])
		}
	}
	return
}

func sanity(filelist []string) error {
	// let's check they're all there..
	for _, v := range filelist {
		_, err := os.Stat(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func usage() {
	fmt.Printf("(S)shd (B)anner (G)enerator v%d\n", sbgVersion)
	fmt.Printf("Usage:\n")
	fmt.Printf("%s [options] <ssh host public key(s)>\n\n", os.Args[0])
	fmt.Printf("where options are:\n")
	flag.PrintDefaults()
	fmt.Printf("\nexample:\n")
	fmt.Printf("\t%s /etc/ssh/ssh_host_*.pub > /etc/ssh/myhostbanner.txt\n", os.Args[0])
}

func main() {
	var keyFileArray [][]string

	bubbleFlag := flag.Bool("b", false, "use bubblebabble digest instead of sha256 (default: sha256)")

	flag.Usage = usage
	flag.Parse()
	argv := flag.Args()

	if len(argv) > 0 {
		err := sanity(argv)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range argv {
			fpbuf, err := sshkf2b(KGCMD, v, *bubbleFlag)
			if err != nil {
				log.Fatal(err)
			}
			keyFileArray = append(keyFileArray, fpbuf)
		}

		printBanner(keyFileArray)
		return
	}
	usage()
	return
}
