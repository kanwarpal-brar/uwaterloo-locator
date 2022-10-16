import ReactMapGL, { Marker, Popup, GeolocateControl } from "react-map-gl";
import washroomData from "../data/Public_Washrooms.json";
import Pin from "./pin";
import { useState } from "react";
import PopupCard from "./PopupCard";

export function MapComponent() {
  const [popupInfo, setPopupInfo] = useState(null);
  return (
    <ReactMapGL
      initialViewState={{
        latitude: 45.387,
        longitude: -75.7,
        zoom: 18,
      }}
      style={{ width: "100vw", height: "100vh" }}
      mapStyle="mapbox://styles/mapbox/light-v9"
      mapboxAccessToken={process.env.REACT_APP_MAPBOX_TOKEN}
    >
      <GeolocateControl position="top-right" />

      {washroomData.features.map((washroom, index) => {
        const latitude = washroom.properties.Y_COORDINATE;
        const longitude = washroom.properties.X_COORDINATE;
        return (
          <Marker
            key={`marker-${index}`}
            longitude={longitude}
            latitude={latitude}
            anchor="bottom"
            onClick={(e) => {
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
          ></PopupCard>
        </Popup>
      )}
    </ReactMapGL>
  );
}

export default MapComponent;
