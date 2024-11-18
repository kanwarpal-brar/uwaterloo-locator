import MapView, { PROVIDER_GOOGLE } from "react-native-maps";
import { UWaterlooRegion } from "../../constants/map-constants";
import { fetchWashroomLocations } from "../../api/location-data-api";
import CustomMapMarker from "../map/custom-marker/custom-marker";
import { useContext, useMemo, useState } from "react";

import {
  StyleProp,
  StyleSheet,
  ViewStyle,
  View,
  StatusBar as NativeStatusBar,
  Platform,
} from "react-native";
import { MapContext, MapContextType, MapModeTypes } from "./map-context";
export type CustomMapProps = {
  style?: StyleProp<ViewStyle>;
};

export default function CustomMap({ style }: CustomMapProps) {
  const washrooms = fetchWashroomLocations();
  const mapContext: MapContextType = useContext(MapContext);
  const [markers, setMarkers] = useState(generateMarkers(mapContext.mode));

  useMemo(() => {
    console.log("Map mode changed to: ", mapContext.mode);
    setMarkers(generateMarkers(mapContext.mode));
  }, [mapContext.mode]);

  function generateMarkers(mapMode: string) {
    switch (mapMode) {
      case MapModeTypes.standard:
        return washrooms.map((washroom, index) => {
          return <CustomMapMarker key={index} location={washroom} />;
        });
      case MapModeTypes.manual:
        return [<CustomMapMarker key={0} location={washrooms[0]} />];
    }
  }

  return (
    <View style={styles.container}>
      <MapView
        style={style ? style : styles.map}
        initialRegion={UWaterlooRegion}
        provider={Platform.OS === "android" ? PROVIDER_GOOGLE : undefined}
        showsMyLocationButton={true}
        showsIndoors={true}
        showsIndoorLevelPicker={false}
        onIndoorBuildingFocused={() => console.log("bruh")}
        showsUserLocation={true}
        userLocationPriority="high"
        showsPointsOfInterest={false}
        loadingEnabled={true}
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
