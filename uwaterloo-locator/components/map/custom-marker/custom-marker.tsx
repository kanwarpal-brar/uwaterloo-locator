import React from "react";
import { Marker } from "react-native-maps";
import { NamedLatLng } from "../../../api/location-data-api";
import CustomPin, { CustomIconType } from "../custom-pin/custom-pin";

type CustomMapMarkerProps = {
  location: NamedLatLng;
  fill?: string;
  icon?: CustomIconType;
};

function CustomMapMarker({
  location,
  fill = "#292929",
  icon = "toilet",
}: CustomMapMarkerProps) {
  return (
    <Marker
      coordinate={{
        latitude: location.latitude,
        longitude: location.longitude,
      }}
      title={location.name}
      tracksViewChanges={false}
    >
      <CustomPin icon={icon} fill={fill} stroke="black" strokeWidth={0.5} />
    </Marker>
  );
}

export default React.memo(CustomMapMarker);
