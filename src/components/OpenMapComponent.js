// import '../styles/OpenMapComponent.css'
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet'
import '../styles/OpenMapComponent.css'
import 'leaflet/dist/leaflet.css';
import washroomData from '../data/Public_Washrooms.json'
export default function OpenMapComponent() {
  return (
    <MapContainer className='.open-map-overlay' center={[43.470, -80.542]} zoom={15}>
        <TileLayer
					attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a>'
    			url="https://tile.openstreetmap.org/{z}/{x}/{y}.png"/>
    </MapContainer>
  )
}
