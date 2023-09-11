import ReactMapGL, {Marker, Popup, GeolocateControl} from "react-map-gl"
import washroomData from '../data/Public_Washrooms.json'
import Pin from './pin';
import { useState } from 'react';
import '../styles/MapBoxComponent.css'
import setCoords from './Card'
import PopupCard from "./PopupCard";

export function MapBoxComponent(){
    const [popupInfo, setPopupInfo] = useState(null);
    // const [coords, setCoords] = useState([]);

    return (
        <ReactMapGL
        initialViewState={{
          longitude: -122.4,
          latitude: 37.8,
          zoom: 18
        }}
        style={{width: "100vw", height: "100vh"}}
        mapStyle="mapbox://styles/mapbox/light-v9"
        mapboxAccessToken={process.env.REACT_APP_MAPBOX_TOKEN}
        >

        <GeolocateControl position="top-right" onGeolocate={(e)=> {setCoords([e.coords.latitude, e.coords.longitude])}}/>
  
          {washroomData.features.map((washroom,index) => {
            const latitude = washroom.properties.Y_COORDINATE;
            const longitude = washroom.properties.X_COORDINATE;
            return (
          <Marker
            key={`marker-${index}`}
            longitude={longitude}
            latitude={latitude}
            anchor="bottom"
            onClick={e => {
              // If we let the click event propagates to the map, it will immediately close the popup
              // with `closeOnClick: true`
              e.originalEvent.stopPropagation();
              setPopupInfo(washroom);

            }}
          >
            <Pin />
          </Marker>
            );
          })}
  
    {popupInfo && (
        <Popup
          anchor="bottom"
          longitude={Number(popupInfo.properties.X_COORDINATE)}
          latitude={Number(popupInfo.properties.Y_COORDINATE)}
          onClose={() => setPopupInfo(null)}
          closeButton={false}
          offset={24}
          style={{
            margin: 0,
            padding: 0,
            boxShadow: "",
          }}
          maxWidth="360px"
        >
          <PopupCard
            name={popupInfo.properties.NAME}
            address={popupInfo.properties.ADDRESS}
            times={popupInfo.properties.HOURS_MONDAY_OPEN + " - " + popupInfo.properties.HOURS_MONDAY_CLOSED}
            f={popupInfo.properties.FAMILY_TOILET}
            d={popupInfo.properties.SPECIAL_TOILET_TYPE}
            c={popupInfo.properties.CHANGE_STATION_CHILD}
          ></PopupCard>
        </Popup>
      )}
      </ReactMapGL>
    )
}

export default MapBoxComponent;