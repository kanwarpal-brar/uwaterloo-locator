import React, { useImperativeHandle, forwardRef, useRef } from "react";
import { MapMarker } from "react-native-maps";
import { NamedLatLng } from "../../../api/location-data-api";
import { FontAwesome6 } from "@expo/vector-icons";

const CustomMapMarker = forwardRef(
  ({ location }: { location: NamedLatLng }, ref) => {
    const markerRef = useRef<MapMarker>(null);

    useImperativeHandle(ref, () => ({
      redraw() {
        console.log("redrawing marker");
        if (markerRef.current) {
          markerRef.current.redraw();
        }
      },
    }));

    return (
      <MapMarker
        ref={markerRef}
        coordinate={{
          latitude: location.latitude,
          longitude: location.longitude,
        }}
        title={location.name}
        tracksViewChanges={false}
      >
        <FontAwesome6
          name="toilet"
          size={15}
          color="#000000"
          style={{
            alignSelf: "center",
          }}
        />
      </MapMarker>
    );
  },
);

export default React.memo(CustomMapMarker);
