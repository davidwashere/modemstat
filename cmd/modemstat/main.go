package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

const statusPath = "DocsisStatus.htm"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	address := "http://192.168.100.1/"
	u, err := url.Parse(address)

	jar, err := cookiejar.New(&cookiejar.Options{})
	check(err)

	client := &http.Client{
		Jar: jar,
	}

	resp, err := client.Get(u.String())
	check(err)

	// fmt.Println(resp)

	buf := new(bytes.Buffer)

	req, err := http.NewRequest(http.MethodGet, u.String()+statusPath, buf)
	req.SetBasicAuth("admin", "password")
	check(err)

	resp, err = client.Do(req)
	check(err)

	// fmt.Println(resp)
	body, err := ioutil.ReadAll(resp.Body)

	bodyS := strings.Split(string(body), "\n")

	var downChans []DownstreamChannel
	var upChans []UpstreamChannel

	inComment := false

	inChannelDown := false
	inChannelUp := false
	channelDownSlug := "InitDsTableTagValue"
	channelUpSlug := "InitUsTableTagValue"

	for _, line := range bodyS {
		if strings.Contains(line, channelDownSlug) {
			inChannelDown = true
			continue
		}

		if strings.Contains(line, channelUpSlug) {
			inChannelUp = true
			continue
		}

		if !inChannelDown && !inChannelUp {
			continue
		}

		tr := strings.TrimSpace(line)

		if strings.HasPrefix(tr, "/*") {
			inComment = true
		}

		if strings.HasPrefix(tr, "*/") {
			inComment = false
		}

		if inComment || !strings.Contains(line, "var tagValueList") {
			continue
		}

		if inChannelDown {
			downChans = parseDownChannels(tr)
			inChannelDown = false
			continue
		}

		if inChannelUp {
			upChans = parseUpChannels(tr)
			inChannelUp = false
			continue
		}
	}

	fmt.Println("\nDownstream Bonded Channels")
	fmt.Println("==========================")
	for _, dc := range downChans {
		fmt.Printf("%2v %7v %7v %13v %5v dBmV %5v dB %12v %12v\n", dc.id, dc.lockStatus, dc.modulation, dc.freq, dc.power, dc.snr, dc.correctables, dc.uncorrectables)
	}

	fmt.Println("\nUpstream Bonded Channels")
	fmt.Println("==========================")
	for _, uc := range upChans {
		fmt.Printf("%2v %7v %6v %5v Ksym/sec %12v %5v dBmV\n", uc.id, uc.lockStatus, uc.channelType, uc.symbolRate, uc.freq, uc.power)
	}

	// fmt.Println()
	// fmt.Printf("%+v\n", downChans)
	// fmt.Println()
	// fmt.Printf("%+v\n", upChans)
}

func parseUpChannels(line string) []UpstreamChannel {
	start := strings.Index(line, "'") + 1
	end := strings.LastIndex(line, "'")

	channelsRaw := line[start:end]
	channelsS := strings.Split(channelsRaw, "|")
	channelsS = channelsS[1:]

	upChans := []UpstreamChannel{}
	upChan := UpstreamChannel{}

	for i, c := range channelsS {
		if i%7 == 0 && i > 0 {
			upChans = append(upChans, upChan)
		}

		adj := i % 7

		c = strings.TrimSpace(c)

		if adj == 1 {
			upChan.lockStatus = c
		} else if adj == 2 {
			upChan.channelType = c
		} else if adj == 3 {
			upChan.id = c
		} else if adj == 4 {
			upChan.symbolRate = c
		} else if adj == 5 {
			upChan.freq = c
		} else if adj == 6 {
			upChan.power = c
		}
	}

	return upChans
}

func parseDownChannels(line string) []DownstreamChannel {
	start := strings.Index(line, "'") + 1
	end := strings.LastIndex(line, "'")

	channelsRaw := line[start:end]
	channelsS := strings.Split(channelsRaw, "|")
	channelsS = channelsS[1:]

	downChannels := []DownstreamChannel{}
	downChannel := DownstreamChannel{}

	for i, c := range channelsS {
		if i%9 == 0 && i > 0 {
			downChannels = append(downChannels, downChannel)
		}

		adj := i % 9

		c = strings.TrimSpace(c)

		if adj == 1 {
			downChannel.lockStatus = c
		} else if adj == 2 {
			downChannel.modulation = c
		} else if adj == 3 {
			downChannel.id = c
		} else if adj == 4 {
			downChannel.freq = c
		} else if adj == 5 {
			downChannel.power = c
		} else if adj == 6 {
			downChannel.snr = c
		} else if adj == 7 {
			downChannel.correctables = c
		} else if adj == 8 {
			downChannel.uncorrectables = c
		}
	}

	return downChannels
}

type DownstreamChannel struct {
	id             string
	lockStatus     string
	modulation     string
	freq           string
	power          string
	snr            string
	correctables   string
	uncorrectables string
}

type UpstreamChannel struct {
	id          string
	lockStatus  string
	channelType string
	symbolRate  string
	freq        string
	power       string
}
