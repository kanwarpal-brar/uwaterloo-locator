import { Region, Camera } from "react-native-maps"

export const UWaterlooRegion: Region = {
    latitude: 43.47198,
    latitudeDelta: 0.012,
    longitude: -80.54396,
    longitudeDelta: 0.012
}

export const UWaterlooInitialCamera: Camera = {
    center: {
        latitude: UWaterlooRegion.latitude,
        longitude: UWaterlooRegion.longitude
    },
    heading: 0,
    pitch: 0,
    zoom: 15.5,
    altitude: 20
}