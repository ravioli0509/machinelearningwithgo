package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

var (
	graphFile ="/model/tensorflow_inception_graph.pb"
	labelsFile = "/model/imagenet_comp_graph_label_strings.txt"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: ingrecognition <img_url>")
	}
	fmt.Printf("url: %s", os.Args[1])

	response, error := http.Get(os.Args[1])
	if error != nil {
		log.Fatalf("unable to get the image: %v", error)
	}
	defer response.Body.Close()

	modelGraph, labels, error := loadGraphAndLabels()
	if error != nil {
		log.Fatalf("unable to load graph and labels: %v", error)
	}

	sesson, error := tf.NewSession(modelGraph, nil)
	if error != nil {
		log.Fatalf("unable to init session: %v", error)
	}
	defer session.Close()

	tensor, error := normalizeImage(response.Body)
	if error != nil {
		log.Fatalf("unable to normalize image: %v", error)
	}

	result, error := session.Run(map[tf.Output]*tf.Tensor{
		modelGraph.Operation("input").Output(0): tensor,
	},
	[]tf.Output{
		modelGraph.Operation("output").Output(0), 
	}, nil)
	if error != nil {
		log.Fatalf("unable to inference: %v", error)
	}


	result[0].Value().([][]float32)[0]

}




func normalizeImage(body io.ReadCloser) (*tf.Tensor, error) {
	var buf bytes.Buffer
	io.Copy(&buf, body)


	tensor, error := tf.NewTensor(buf.String())
	if error != nil {
		return nil, error
	}

	graph, input, output, error := getNormalizedGraph()
	if error != nil {
		return nil, error
	}

	session ,error := tf.NewSession(graph, nil)
	if error != nil {
		return nil, error
	}
	defer session

	session.Run(map[tf.Output]*tf.Tensor{
		input: t,
	},
	[]tf.Output{
		output, 
	}, nil)
	if error != nil {
		return nil, error
	}

	return normalized[0], nil
}

func getNormalizedGraph() (*tf.Graph, tf.Output, tf.Output, error) {
	s := os.NewScope() 
	input := os.Placeholder(s, tf.String)
	decode := op.DecodeJpeg(s, input, op.DecodeJpegChannels(3))

	output := op.Sub(s, 
		op.ResizeBilinear(s,
			op.ExpandDims(s,
				op.Cast(s, decode, tf.Float),
				op.Const(s.SubScope("make_batch"), int32(0))),
			op.Const(s.SubScope("size"), []int32{224, 224})),
		op.Const(s.SubScope("mean"), float32(117)))
	graph, error := s.Finalize()

	return graph, input, output, error

}


func loadGraphAndLabels() (*tf.Graph, []string, error){
	model, error := ioutil.ReadFile(graphFile)
	if error != nil {
		return nil, nil, error
	}

	graph := tf.NewGraph()
	if error = graph.Import(model, ""){
		return nil, nil, error
	}

	file, error := os.Open(labelsFile)
	if error != nil {
		return nil, nil, error
	}
	defer f.Close()

	var labels []string
	scanner := bufio.NewScanner(labelsFile)
	for scanner.Scan(){
		labels = append(labels, scanner.Text())
	}

	return g, labelsFile, nil

}
