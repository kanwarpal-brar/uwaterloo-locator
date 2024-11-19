import React, { useContext, useEffect, useRef, useState, useMemo } from "react";
import {
  StyleProp,
  StyleSheet,
  ViewStyle,
  View,
  StatusBar as NativeStatusBar,
} from "react-native";
import MapView, { MapMarker, PROVIDER_GOOGLE } from "react-native-maps";
import { UWaterlooRegion } from "../../constants/map-constants";
import { fetchWashroomLocations } from "../../api/location-data-api";
import { MapContext, MapContextType, MapModeTypes } from "./map-context";
import CustomMarkerV2, {
  CustomMarkerRef,
} from "./custom-marker/custom-marker-v2";

export type CustomMapProps = {
  style?: StyleProp<ViewStyle>;
};

export default function CustomMap({ style }: CustomMapProps) {
  const mapRef = useRef<MapView>(null);
  const markerRefs = useRef<CustomMarkerRef[]>([]);
  const washrooms = fetchWashroomLocations();
  const mapContext: MapContextType = useContext(MapContext);
  const [markers, setMarkers] = useState<JSX.Element[]>([]);

  // Initialize markers based on map mode
  useMemo(() => {
    setMarkers(generateMarkers(mapContext.mode));
  }, [mapContext.mode]);

  // Function to redraw all markers
  const redrawAllMarkers = () => {
    console.log("map redraw");
    markerRefs.current.forEach((ref) => {
      if (ref) {
        ref.redraw();
      }
    });
  };

  // Function to redraw specific marker
  const redrawMarker = (index: number) => {
    if (markerRefs.current[index]) {
      markerRefs.current[index].redraw();
    }
  };

  // Handler for map load
  const handleMapLoad = () => {
    redrawAllMarkers();
  };

  useEffect(() => {
    redrawAllMarkers();
  }, [mapContext.mode]);

  function generateMarkers(mapMode: string): JSX.Element[] {
    switch (mapMode) {
      case MapModeTypes.standard:
        return washrooms.map((washroom, index) => (
          <CustomMarkerV2
            key={index}
            ref={(ref) => {
              if (ref) {
                console.log(`set ref ${index}`);
                markerRefs.current[index] = ref;
              }
            }}
            coordinate={washroom}
          />
        ));
      case MapModeTypes.manual:
        return [
          <CustomMarkerV2
            key={0}
            ref={(ref) => {
              if (ref) {
                console.log(`set ref ${0}`);
                markerRefs.current[0] = ref;
              }
            }}
            coordinate={washrooms[0]}
          />,
        ];
      default:
        return [];
    }
  }

  return (
    <View style={styles.container}>
      <MapView
        ref={mapRef}
        style={style ? style : styles.map}
        initialRegion={UWaterlooRegion}
        provider={PROVIDER_GOOGLE}
        showsMyLocationButton={true}
        showsIndoors={true}
        showsIndoorLevelPicker={false}
        onIndoorBuildingFocused={() => console.log("bruh")}
        onMapLoaded={handleMapLoad}
      >
        {markers}
      </MapView>
    </View>
  );
}

const styles = StyleSheet.create({
  map: {
    ...StyleSheet.absoluteFillObject,
    marginTop: NativeStatusBar.currentHeight
      ? NativeStatusBar.currentHeight + 2
      : 32,
    zIndex: 0,
  },
  container: {
    ...StyleSheet.absoluteFillObject,
    zIndex: 0,
    flexDirection: "row-reverse",
  },
});
