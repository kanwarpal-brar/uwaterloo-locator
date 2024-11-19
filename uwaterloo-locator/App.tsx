import { StatusBar } from "expo-status-bar";
import { StyleSheet, View } from "react-native";
import CustomMap from "./components/map/custom-map";
import Footer from "./components/home-ui/footer";
import { MapDataProvider } from "./components/map/map-context";

export default function App() {
  return (
    <View style={styles.container}>
      <MapDataProvider>
        <CustomMap />
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
