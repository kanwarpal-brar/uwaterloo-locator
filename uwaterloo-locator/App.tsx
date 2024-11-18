import { StatusBar } from "expo-status-bar";
import { StyleSheet, Text, View } from "react-native";
import CustomMap from "./components/map/custom-map";
import Footer from "./components/home-ui/footer";
import { MapDataProvider } from "./components/map/map-context";
import * as Location from "expo-location";
import { useEffect, useMemo, useState } from "react";

export default function App() {
  const [location, setLocation] = useState(null as any);
  const [errorMsg, setErrorMsg] = useState(null as any);
  const [haveLocationPerm, setHaveLocationPerm] = useState(false);

  useMemo(() => {
    (async () => {
      let { status } = await Location.requestForegroundPermissionsAsync();
      if (status !== "granted") {
        setErrorMsg("Permission to access location was denied");
        return;
      }
      setHaveLocationPerm(true);

      let location = await Location.getCurrentPositionAsync({});
      setLocation(location);
    })();
  }, []);

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
