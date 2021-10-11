import React, {useEffect, useState} from 'react';
import { Route, Switch, BrowserRouter } from 'react-router-dom';
import axios, {AxiosResponse} from 'axios';
import {FaHeart} from "react-icons/all";

const Names =["Vladi","Eliran","Yuval","Yaron","Eli","Yossi","Ben","Daniel","Iris","Bat-Hen",
            "Ran","Dolev","Shlaev","Boaz","Michael","Dana","Olga","Yafit","Mor","Ilay"];
const Cities =["Tel Aviv","Herzliya","Netanya","Rishpon","Ramat-Gan","Givatayim","Ramat-Sharon","Karkur","Hadera","Kfar-Saba"
            ,"Beer-Sheva","Eilat","Dimona","Jerusalem","Modiin","Shoham","Caesarea","Haifa","Akko","Atlit"];

const RandomName=()=> {
    const rand =  Math.floor(Math.random()* 20)
    return Names[rand]
}

const RandomCities=()=> {
    const rand = Math.floor(Math.random()*20)
    return Cities[rand]
}

const API= axios.create({
    baseURL: "http://localhost:8080",
    headers: {
        "Content-type": "application/json"
    }
})

function App() {
    return (
        <BrowserRouter>
            <Switch>
                <Route exact path='/'>
                    <Home />
                </Route>
            </Switch>
        </BrowserRouter>
    );
}
const Home = () => {

    return (
            <div>
            <NavBar />
            <header className="text-center text-white bg-primary masthead">
                <div className="container" >
                    <img className="img-fluid d-block mx-auto mb-3" src="assets/img/taxi.png" alt={"profile picture"} width="300" height="300"></img>
                    <h1>GetMonit</h1>
                    <hr className="star-light"></hr>
                        <h2 className="font-weight-light mb-0">The smartest way to move around in cities.</h2>
                        <h2 className="font-weight-light mb-0">Forget expensive taxi rides or slow public transport.</h2>
                         <h2 className="font-weight-light mb-0">Get a ride in minutes</h2>
                </div>
            </header>
            <div className="d-flex flex-column" id="content-wrapper">
                <div id="content">
                    <div className="container-fluid">
                        <div className="row">
                            <div className="col-lg-7 col-xl-7">
                                <div className="card shadow mb-4">
                                    <div className="card-header d-flex justify-content-between align-items-center">
                                        <Rides />
                                        </div>
                                    </div>
                                </div>
                            <div className="col-lg-5 col-xl-5">
                                <div className="card shadow mb-4">
                                    <div className="card-header d-flex justify-content-between align-items-center">
                                        <div className="card-body">
                                            <MessageList />
                                        </div>
                                    </div>

                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            </div>
    );
};


const Rides = () => {
    const [name,setName]=useState("")
    const [source,setSource]=useState("")
    const [destination,setDestination]=useState("")
    const [message,SetMessage]=useState("")

    const NewTravel=async(n:string,s:string,d:string)=>{
        await API.post('/travel',
            {name:n,source:s,destination:d})
            .then(res=>{if (res.status==200){
                SetMessage("Travel request was recorded ,We are looking for a free drive, This may take several minutes...")
                }})
            .catch(err=>SetMessage(err))
    }

    const GenerateTravels=()=>{
            for(let i=0;i<10;i++){
                const n:string=RandomName()
                const s:string=RandomCities()
                const d:string=RandomCities()
                NewTravel(n,s,d)
            }
    }
    return (
            <div className="container" >
                <h2 className="text-uppercase text-center text-secondary mb-0">Get A Ride</h2>
                <hr className="star-dark mb-2"></hr>
                <div className="row">
                    <div className="col-lg-8 mx-auto">
                        <form id="travelForm" name="createTravel" noValidate={false}>
                            <div className="control-group">
                                <div className="mb-0 form-floating controls pb-2"><input className="form-control" type="text" id="name" required={true} placeholder="Name" onChange={(e)=>{setName(e.target.value)}}></input><label className="form-label">Name</label><small
                                    className="form-text text-danger help-block"></small></div>
                            </div>
                            <div className="control-group">
                                <div className="mb-0 form-floating controls pb-2"><input className="form-control" type="source" id="source" required={true} placeholder="Source" onChange={(e)=>{setSource(e.target.value)}}></input>
                                    <label className="form-label">Source</label><small className="form-text text-danger help-block"></small>
                                </div>
                            </div>
                            <div className="control-group">
                                <div className="mb-0 form-floating controls pb-2"><input className="form-control" type="tel" id="Destination" required={true} placeholder="Destination" onChange={(e)=>{setDestination(e.target.value)}}></input>
                                    <label className="form-label">Destination</label><small className="form-text text-danger help-block"></small>
                                </div>
                            </div>
                            <div id="success"></div>
                            <div>
                                <button className="btn btn-primary btn-xl" type="button" disabled={name==="" || source==="" || destination===""?true:false} onClick={()=>NewTravel(name,source,destination)}>Find Taxi</button>
                                <button className="btn btn-primary btn-xl" type="button" onClick={()=>GenerateTravels()}>Generate Travels
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
    );
};

const NavBar = () => {
    return (
        <nav className="navbar navbar-light navbar-expand-lg fixed-top bg-secondary text-uppercase" id="mainNav">
            <div className="container"><a className="navbar-brand" href="#page-top">GetMonit</a>
                <button data-bs-toggle="collapse" data-bs-target="#navbarResponsive" className="navbar-toggler text-white bg-primary navbar-toggler-right text-uppercase rounded"
                        aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation"><i className="fa fa-bars"></i></button>
                <div className="collapse navbar-collapse" id="navbarResponsive">
                    <ul className="navbar-nav ms-auto">
                        <li className="nav-item mx-0 mx-lg-1"><a className="nav-link py-3 px-0 px-lg-3 rounded" href="/">I</a></li>
                        <li className="nav-item mx-0 mx-lg-1"><a className="nav-link py-3 px-0 px-lg-3 rounded" href="/"><FaHeart/></a></li>
                        <li className="nav-item mx-0 mx-lg-1"><a className="nav-link py-3 px-0 px-lg-3 rounded" href="/">SAP</a></li>
                    </ul>
                </div>
            </div>
        </nav>
    );
};


const MessageList = () => {
    const [messages,SetMessages]=useState<message[]>([])

    const ClearHandler=async()=>{
        await API.put('/clear')
            .then(res=> res.data)
            .catch(err=>console.log(err))

        SetMessages([])
    }
    useEffect(()=>{
        const interval=setInterval(async()=>{
            await API.get('/getnotifications')
                .then(res=> SetMessages(res.data as message[]))
                .catch(err=>console.log(err))
        },7000);

        return ()=>{clearInterval(interval)};
    },[])


    return (
        <div className="card shadow mb-4">
            <div className="card-header py-3">
                <h6 className="text-primary fw-bold m-0">Messages</h6>
                <button className="btn btn-primary btn-s" type="button" onClick={()=>ClearHandler()}>Clear</button>
            </div>
            <ul className="list-group list-group-flush">
                {messages.reverse().map(m=>(
                    <div>
                        <Message data={m}/>
                    </div>
                    )
                )}
            </ul>
        </div>
    );
};
interface message{
    taxiId:string
    message:string
}
interface MessageProps {
    data:message
}
const Message = (Props:MessageProps) => {
    return (
        <li className="list-group-item">
            <div className="row align-items-center no-gutters">
                <div className="col me-2">
                    <h6 className="mb-0"><strong>Taxi : {Props.data.taxiId}</strong></h6><span
                    className="text-xs">{Props.data.message}</span>
                </div>
            </div>
        </li>
    );
};
export default App;
