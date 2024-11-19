import React from "react";
import { Marker } from "react-native-maps";
import { NamedLatLng } from "../../../api/location-data-api";
import CustomPin from "../custom-pin/custom-pin";

function CustomMapMarker({ location }: { location: NamedLatLng }) {
  return (
    <Marker
      coordinate={{
        latitude: location.latitude,
        longitude: location.longitude,
      }}
      title={location.name}
      tracksViewChanges={false}
    >
      <CustomPin
        icon={"toilet"}
        fill="#292929"
        stroke="black"
        strokeWidth={0.5}
      />
    </Marker>
  );
}

export default React.memo(CustomMapMarker);
