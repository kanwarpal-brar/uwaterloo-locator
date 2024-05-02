import { Marker } from "react-native-maps";
import { NamedLatLng } from "../../../api/location-data-api";
import CustomPin from "../custom-pin/custom-pin";
import { FontAwesome6 } from "@expo/vector-icons";

export default function CustomMapMarker({
  location,
}: {
  location: NamedLatLng;
}) {
  const toiletIcon = (
    <FontAwesome6
      name="toilet"
      size={15}
      color="#ededed"
      style={{ alignSelf: "center" }}
    />
  );
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
        icon={"person-half-dress"}
        fill="#292929"
        stroke="black"
        strokeWidth={0.5}
      />
    </Marker>
  );
}
