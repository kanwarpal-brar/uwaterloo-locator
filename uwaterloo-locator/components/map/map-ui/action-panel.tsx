import { useContext } from "react";
import { View, Text, StyleSheet, Pressable } from "react-native";
import {
  MapActionTypes,
  MapContext,
  MapDispatchContext,
  MapModeTypes,
} from "../map-context";

export default function ActionPanel() {
  const mapContext = useContext(MapContext);
  const mapDispatchContext = useContext(MapDispatchContext);

  const buttons = [
    {
      name: "Create",
      onPress: () => {
        console.log(mapContext.mode);
        mapDispatchContext({
          type:
            mapContext.mode === MapModeTypes.standard
              ? MapActionTypes.SET_MANUAL_MODE
              : MapActionTypes.SET_STANDARD_MODE,
        });
      },
    },
    {
      name: "View Events",
      onPress: () => console.log("Test"),
    },
  ];
  return (
    <View style={styles.panel}>
      {buttons.map((button, index) => {
        return (
          <Pressable
            key={index}
            style={({ pressed }) => [
              pressed ? actionStyles.button_pressed : styles.button,
            ]}
            onPress={button.onPress}
          >
            <Text>{button.name}</Text>
          </Pressable>
        );
      })}
    </View>
  );
}

const styles = StyleSheet.create({
  panel: {
    flex: 1,
    zIndex: 1,
    alignContent: "center",
    justifyContent: "space-evenly",
    flexDirection: "row",
    paddingTop: "5%",
    paddingBottom: "5%",
  },
  button: {
    aspectRatio: 5 / 3,
    width: "25%",
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "white",
    borderRadius: 10,
    borderColor: "#ededed",
    borderWidth: 1,
    elevation: 2.5, // Add this line
  },
});

const actionStyles = StyleSheet.create({
  button_pressed: {
    ...styles.button,
    backgroundColor: "#ededed",
  },
});
