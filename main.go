package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/velann21/match_todo_gateway_srv/pkg/middleware"
	"github.com/velann21/match_todo_gateway_srv/pkg/routes"
	rmPf "github.com/velann21/todo-commonlib/proto_files/resource_manager"
	userPf "github.com/velann21/todo-commonlib/proto_files/users_srv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/resolver"
	"log"
	"net/http"
	"os"
	"time"
)


func main() {

	//var log = logrus.New()
	//log.Formatter = new(logrus.TextFormatter)
	//log.Formatter.(*logrus.TextFormatter).DisableColors = true
	//log.Formatter.(*logrus.TextFormatter).DisableTimestamp = true
	//if rcl, err := NewRcLogHook("localhost:50052"); err == nil {
	//	log.Hooks.Add(rcl)
	//}
	//log.Out = os.Stdout

	r := mux.NewRouter().StrictSlash(false)
	r.Use(middleware.TraceLogger())
	r.Use(middleware.Authentication())
	r.Use(middleware.Metrics())

	r.Use()

	resolver.SetDefaultScheme("dns")

	mainRoutes := r.PathPrefix("/api/v1").Subrouter()

	usersConn , err := DialUsersSrv()
	if err != nil{
		logrus.Error("Something went wrong while calling USER grpc server")
		os.Exit(1)
	}
	userClient := userPf.NewUserManagementServiceClient(usersConn)
	users := routes.UsersRoutes{UsersServiceGrpcClient: userClient,}
	users.UsersRoutes(mainRoutes)
	go RegisterUsersGrpcConnectionState(usersConn)

	rmConn, err := DialResourceManager()
	if err != nil{
		logrus.Error("Something went wrong while calling RM grpc server")
		os.Exit(1)
	}
	rmClient := rmPf.NewResourceManagerServiceClient(rmConn)
	rmRoutes := routes.ResourceManagerRoutes{ResourceManagerGrpcClient:rmClient, Log: nil}
	rmRoutes.ResourceManagerRoutes(mainRoutes)
	go RegisterRmGrpcConnectionState(rmConn)


	logrus.WithField("EventType", "Bootup").Info("Booting up server at port : " + "8081")
	if err := http.ListenAndServe(":8081", r); err != nil {
		logrus.WithField("EventType", "Server Bootup").WithError(err).Error("Server Bootup Error")
		log.Fatal(err)
		return
	}
}

func DialResourceManager()(*grpc.ClientConn, error){
	rmConn , err := grpc.Dial(
		"localhost:50052",
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	)
	if err != nil{
		logrus.Error("Something went wrong while calling AM grpc server")
		return nil, err
	}
	return rmConn, nil
}

func DialUsersSrv()(*grpc.ClientConn, error){
	rmConn , err := grpc.Dial(
		"localhost:50051",
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	)
	if err != nil{
		logrus.Error("Something went wrong while calling AM grpc server")
		return nil, err
	}
	return rmConn, nil
}

func RegisterRmGrpcConnectionState(conn *grpc.ClientConn) {
	go func() {
		for  {
			state := conn.GetState()
			if (state == connectivity.TransientFailure) || (state == connectivity.Shutdown) {
				fmt.Println("RegisterRmGrpcConnectionState is down")
				time.Sleep(time.Second * 10)
			}
			fmt.Println("State ", state)
			time.Sleep(time.Second * 2)
		}
	}()
}

func RegisterUsersGrpcConnectionState(conn *grpc.ClientConn) {
	go func() {
		for  {
			state := conn.GetState()
			if (state == connectivity.TransientFailure) || (state == connectivity.Shutdown) {
				fmt.Println("RegisterRmGrpcConnectionState is down")
				time.Sleep(time.Second * 10)
			}
			fmt.Println("State ", state)
			time.Sleep(time.Second * 2)
		}
	}()
}





type RclogHook struct {
	clientConn *grpc.ClientConn
}

func NewRcLogHook(address string)(*RclogHook, error){
	rmConn , err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	)
	if err != nil{
		logrus.Error("Something went wrong while calling AM grpc server")
		return nil, err
	}
	fmt.Println("Connection made")
	return &RclogHook{clientConn:rmConn}, nil
}

type Mystruct struct {
	Level logrus.Level `json:"level"`
	Message string `json:"msg"`
	Number int `json:"number"`
	Omg bool `json:"omg"`
}

func (hook *RclogHook) Fire(entry *logrus.Entry) error {
	data := entry.Data
	myst := &Mystruct{}
	byData, err := json.Marshal(data)
	err = json.Unmarshal(byData, &myst)
	if err != nil{
		fmt.Println("error occured Unmarshal",err)
	}
	fmt.Println("Data: ", myst.Number)


	_, err = entry.String()
	if err != nil {
		//fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}


	switch entry.Level {
	case logrus.PanicLevel:
		fmt.Println("Sending the data to panic")

	case logrus.FatalLevel:
		fmt.Println("Sending the data to FatalLevel")
		trans := http.Transport{}
		client := http.Client{
			Transport:&trans,
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/api/v1/users/test", nil)
		if err != nil{

		}
		resp, err := client.Do(req)
		if err != nil{
			log.Fatal("failed")
		}
		fmt.Println(resp.Body)
		fmt.Println(resp.Status)
	case logrus.ErrorLevel:
		fmt.Println("Sending the data to ErrorLevel")
		trans := http.Transport{}
		client := http.Client{
			Transport:&trans,
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/api/v1/users/test", nil)
		if err != nil{

		}
		resp, err := client.Do(req)
		if err != nil{
			log.Fatal("failed")
		}
		fmt.Println(resp.Body)
		fmt.Println(resp.Status)
	case logrus.WarnLevel:
		fmt.Println("Sending the data to WarnLevel")
		trans := http.Transport{}
		client := http.Client{
			Transport:&trans,
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/api/v1/users/test", nil)
		if err != nil{

		}
		resp, err := client.Do(req)
		if err != nil{
			log.Fatal("failed")
		}
		fmt.Println(resp.Body)
		fmt.Println(resp.Status)
	case logrus.InfoLevel:
		fmt.Println("Inside InfoLevel")
		rmClient := rmPf.NewResourceManagerServiceClient(hook.clientConn)
		er := rmPf.EventsRequests{
			EventType:"DeployApp",
			ServiceName:"GatewaySrv",
			Time: time.Now().UTC().String(),
			ActionType:"Request",
			TraceID: "skdjksjd8",
			Transaction:&rmPf.EventsTransaction{
				EventType:"DeployApp",
			},
		}
		eventResponse, err := rmClient.CollectEvent(context.Background(), &er)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println("Final Resp", eventResponse)
	case logrus.DebugLevel, logrus.TraceLevel:
		fmt.Println("Sending the data to DebugLevel, TraceLevel")
		trans := http.Transport{}
		client := http.Client{
			Transport:&trans,
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/api/v1/users/test", nil)
		if err != nil{

		}
		resp, err := client.Do(req)
		if err != nil{
			log.Fatal("failed")
		}
		fmt.Println(resp.Body)
		fmt.Println(resp.Status)
	default:
		return nil
	}
	return nil
}

func (hook *RclogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}