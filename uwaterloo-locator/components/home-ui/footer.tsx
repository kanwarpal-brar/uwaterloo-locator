import { View, Text, StyleSheet } from "react-native";

export default function Footer() {
  return (
    <View style={{ ...styles.footer }}>
      <Text>Footer</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  footer: {
    zIndex: 1,
    backgroundColor: "white",
    borderTopLeftRadius: 20,
    borderTopRightRadius: 20,
    width: "100%",
    height: "12%",
    padding: 10,
  },
});
