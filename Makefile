count := 1000
gogcdcmd := go run cmd/gcd/gcd.go --count=$(count)
pygcdcmd := ./gcd.py count=$(count)
golcmcmd := go run cmd/lcm/lcm.go --count=$(count)
pylcmcmd := ./lcm.py count=$(count)

run:
	go run cmd/main/main.go

gogcd:
	$(gogcdcmd)

pygcd:
	$(pygcdcmd)

pylcm:
	$(pylcmcmd)

golcm:
	$(golcmcmd)

comparegcd:
	$(gogcdcmd) > go.gcd.out
	$(pygcdcmd) > py.gcd.out
	diff go.gcd.out py.gcd.out

comparelcm:
	$(golcmcmd) > go.lcm.out
	$(pylcmcmd) > py.lcm.out
	diff go.lcm.out py.lcm.out
