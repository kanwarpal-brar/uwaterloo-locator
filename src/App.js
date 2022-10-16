import logo from './logo.svg';
import './App.css';
import 'mapbox-gl/dist/mapbox-gl.css';
import ReactMapGL, {Marker, Popup, GeolocateControl} from "react-map-gl"
import washroomData from './data/Public_Washrooms.json'
import Pin from './components/pin';
import MapComponent from './components/MapComponent';
import HStack from './components/HStack';
import { useState } from 'react';
import TStack from './components/TStack';
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
      <MapComponent className="map-component"></MapComponent>
      <TStack/>
      <HStack/>
    </div>
  );
}

export default App;
