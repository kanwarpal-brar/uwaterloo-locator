import { StatusBar } from "expo-status-bar";
import { StyleSheet, View } from "react-native";
import CustomMap from "./components/map/custom-map";
import Footer from "./components/home-ui/footer";
import {
  MapActionTypes,
  MapDataProvider,
  MapDispatchContext,
} from "./components/map/map-context";
import * as Location from "expo-location";
import { useContext, useEffect, useState } from "react";

export default function App() {
  const mapDispatchContext = useContext(MapDispatchContext);
  const [errorMsg, setErrorMsg] = useState(null as any);
  const [haveLocationPerm, setHaveLocationPerm] = useState(false);

  useEffect(() => {
    (async () => {
      const { status } = await Location.requestForegroundPermissionsAsync();
      if (status !== "granted") {
        setErrorMsg("Permission to access location was denied");
        return;
      }
      setHaveLocationPerm(true);

      const location = await Location.getCurrentPositionAsync({});
      mapDispatchContext({
        type: MapActionTypes.SET_USER_LOCATION,
        payload: location,
      });
    })();
  });

  return (
    <View style={styles.container}>
      <MapDataProvider>
        {haveLocationPerm && <CustomMap />}
        <Footer />
        <StatusBar
          style="light"
          hidden={false}
          networkActivityIndicatorVisible={true}
        />
      </MapDataProvider>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    flexDirection: "column",
    backfaceVisibility: "hidden",
    alignItems: "center",
    justifyContent: "flex-end",
    height: "80%",
    backgroundColor: "black",
  },
});
