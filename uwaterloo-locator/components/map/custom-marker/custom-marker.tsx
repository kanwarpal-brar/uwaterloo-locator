import { useEffect, useRef } from "react";
import { MapMarker } from "react-native-maps";
import { NamedLatLng } from "../../../api/location-data-api";
import CustomPin from "../custom-pin/custom-pin";

export default function CustomMapMarker({
  location,
}: {
  location: NamedLatLng;
}) {
  const markerRef = useRef<MapMarker>(null);
  // const [reloadState, setReloadState] = useState(false);

  // const toiletIcon = (
  //   <FontAwesome6
  //     name="toilet"
  //     size={15}
  //     color="#ededed"
  //     style={{ alignSelf: "center" }}
  //   />
  // );

  const redraw = () => {
    markerRef.current?.redraw();
  };

  // useEffect(() => {
  //   console.log("mounted");
  //   redraw();
  // });

  return (
    <MapMarker
      ref={markerRef}
      coordinate={{
        latitude: location.latitude,
        longitude: location.longitude,
      }}
      title={location.name}
      // onLayout={() => setReloadState(true)}
      tracksViewChanges={false}
    >
      <CustomPin
        icon={"toilet"}
        fill="#292929"
        stroke="black"
        strokeWidth={0.5}
      />
    </MapMarker>
  );
}
