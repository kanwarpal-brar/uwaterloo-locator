import { StatusBar } from 'expo-status-bar';
import { StatusBar as NativeStatusBar } from 'react-native';
import { StyleSheet, Text, View, } from 'react-native';
import MapView, { PROVIDER_GOOGLE } from 'react-native-maps';
import { UWaterlooRegion } from './constants/map-constants';
import { NamedLatLng, fetchWashroomLocations } from './api/location-data-api'
import CustomMapMarker from './components/map/custom-marker/custom-marker';




export default function App() {
  const washrooms = fetchWashroomLocations();
  return (
    <View style={styles.container} >
      <MapView
        style={styles.map}
        region={UWaterlooRegion}
        provider={PROVIDER_GOOGLE}
        showsMyLocationButton={true}
        showsIndoors={true}
        showsIndoorLevelPicker={false}
        onIndoorBuildingFocused={() => console.log('bruh')}
      >
        {
          washrooms.map(
            (washroom, index) => {
              return <CustomMapMarker
              key={index}
              location={washroom}
              />
            }
          )
        }
      </MapView>

      <StatusBar
        style="light"
        hidden={false}
        networkActivityIndicatorVisible={true}
      />
      
      <View style={{...styles.footer}}>
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
    marginTop: NativeStatusBar.currentHeight ? NativeStatusBar.currentHeight + 2 : 32,
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
