package main

type DataPuller interface {
	PullData()
}

type Puller struct {
}

func (p *Puller) PullData() {

}
