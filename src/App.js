import logo from './logo.svg';
import './App.css';
import 'mapbox-gl/dist/mapbox-gl.css';
import ReactMapGL, {Marker, Popup, GeolocateControl} from "react-map-gl"
import washroomData from './data/Public_Washrooms.json'
import Pin from './components/pin';
import MapBoxComponent from './components/MapBoxComponent';
import HStack from './components/HStack';
import { useState } from 'react';
import TStack from './components/TStack';
import OpenMapComponent from './components/OpenMapComponent';
function App() {

  // TODO: Viewport useState
//   const [viewport, setViewport] = useState({
//     latitude: 45.4211,
//     longitude: -75.6903,
//     width: "100vw",
//     height: "100vh",
//     zoom: 10
//   });


// TODO Generate map after we know the location

  return (
    <div className="App">
      <OpenMapComponent ></OpenMapComponent>
      <TStack/>
      <HStack/>
    </div>
  );
}

export default App;
