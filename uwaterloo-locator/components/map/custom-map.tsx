import MapView, { PROVIDER_GOOGLE } from "react-native-maps";
import { UWaterlooRegion } from "../../constants/map-constants";
import { fetchWashroomLocations } from "../../api/location-data-api";
import CustomMapMarker from "../map/custom-marker/custom-marker";

import {
  StyleProp,
  StyleSheet,
  ViewStyle,
  View,
  StatusBar as NativeStatusBar,
} from "react-native";

export type CustomMapProps = {
  style?: StyleProp<ViewStyle>;
};

export default function CustomMap({ style }: CustomMapProps) {
  const washrooms = fetchWashroomLocations();
  return (
    <View style={styles.container}>
      <MapView
        style={style ? style : styles.map}
        region={UWaterlooRegion}
        provider={PROVIDER_GOOGLE}
        showsMyLocationButton={true}
        showsIndoors={true}
        showsIndoorLevelPicker={false}
        onIndoorBuildingFocused={() => console.log("bruh")}
      >
        {washrooms.map((washroom, index) => {
          return <CustomMapMarker key={index} location={washroom} />;
        })}
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
