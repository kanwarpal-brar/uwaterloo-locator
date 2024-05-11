import { StatusBar } from "expo-status-bar";
import { StyleSheet, Text, View } from "react-native";
import CustomMap from "./components/map/custom-map";
import Footer from "./components/home-ui/footer";

export default function App() {
  return (
    <View style={styles.container}>
      <CustomMap />
      <Footer />
      <StatusBar
        style="light"
        hidden={false}
        networkActivityIndicatorVisible={true}
      />
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
