import { StatusBar } from 'expo-status-bar';
import { StatusBar as NativeStatusBar } from 'react-native';
import { StyleSheet, Text, View, } from 'react-native';
import MapView from 'react-native-maps';

export default function App() {
  return (
    <View style={styles.container} >
      <MapView style={styles.map} region={{latitude: 43.47198, latitudeDelta: 0, longitude: -80.54396, , longitudeDelta:}}/>
      <StatusBar style="light" hidden={false} networkActivityIndicatorVisible={true}/>
      <View style={styles.footer}>
        <Text>Header</Text>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    flexDirection: 'column',
    backfaceVisibility: 'hidden',
    alignItems: 'center',
    justifyContent: 'flex-end',
    height: "80%",
    backgroundColor: 'black'
  },
  map: {
    ...StyleSheet.absoluteFillObject,
    marginTop: NativeStatusBar.currentHeight ? NativeStatusBar.currentHeight + 2 : 30,
    zIndex: 0,
  },
  footer: {
    zIndex: 1,
    backgroundColor: 'white',
    borderTopLeftRadius: 20,
    borderTopRightRadius: 20,
    width: '100%',
    height: '12%',
    padding: 10
  }
});
