import MapView, { PROVIDER_GOOGLE } from "react-native-maps";
import { UWaterlooRegion } from "../../constants/map-constants";
import { fetchWashroomLocations } from "../../api/location-data-api";
import CustomMapMarker from "../map/custom-marker/custom-marker";
import { useContext, useEffect, useMemo, useState } from "react";
import {
  StyleProp,
  StyleSheet,
  ViewStyle,
  View,
  StatusBar as NativeStatusBar,
  Platform,
} from "react-native";
import {
  MapActionTypes,
  MapContext,
  MapContextType,
  MapDispatchContext,
  MapModeTypes,
} from "./map-context";
import * as Location from "expo-location";

export type CustomMapProps = {
  style?: StyleProp<ViewStyle>;
};

export default function CustomMap({ style }: CustomMapProps) {
  const [haveLocationPerm, setHaveLocationPerm] = useState(false);
  const mapDispatchContext = useContext(MapDispatchContext);
  const mapContext: MapContextType = useContext(MapContext);
  const washrooms = fetchWashroomLocations();
  const [markers, setMarkers] = useState(generateMarkers(mapContext.mode));
  const [lastUserLocMarker, setLastUserLocMarker] =
    useState<JSX.Element | null>(null);
  const initialRegion = UWaterlooRegion;

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

  useEffect(() => {
    // on mount
    (async () => {
      console.log("running");
      const { status } = await Location.requestForegroundPermissionsAsync();
      if (status !== "granted") {
        return;
      }
      const location = await Location.getCurrentPositionAsync({});

      // Set initialRegion to user's location
      initialRegion.latitude = location.coords.latitude;
      initialRegion.longitude = location.coords.longitude;

      // Trigger Map
      setHaveLocationPerm(true);

      // update context with user's location
      mapDispatchContext({
        type: MapActionTypes.SET_USER_LOCATION,
        payload: location,
      });
    })();
  }, []);

  useMemo(() => {
    if (!mapContext.lastUserLocation) return;
    setLastUserLocMarker(
      <CustomMapMarker
        key={washrooms.length + 1}
        icon="person"
        fill="red"
        location={{
          name: "Last User Location",
          latitude: mapContext.lastUserLocation?.coords.latitude,
          longitude: mapContext.lastUserLocation?.coords.longitude,
        }}
      />,
    );
  }, [mapContext.lastUserLocation]);

  return (
    <View style={styles.container}>
      {haveLocationPerm && (
        <MapView
          style={style ? style : styles.map}
          initialRegion={initialRegion}
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
          {/* {lastUserLocMarker} */}
        </MapView>
      )}
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
