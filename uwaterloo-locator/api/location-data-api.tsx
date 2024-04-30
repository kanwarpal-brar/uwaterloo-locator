import { LatLng } from "react-native-maps";
import buildings from "../constants/buildings.json";

export type NamedLatLng = {
    name: string;
    latitude: number;
    longitude: number;
};

export function fetchWashroomLocations(): Array<NamedLatLng> {
    // TODO this is currently mocked by parsing a geojson: connect to backend later
    return buildings.features.map(
        build =>  {
            return {
                name: build.properties.buildingName,
                latitude: build.properties.latitude || 0,
                longitude: build.properties.longitude || 0
            }
        }
    );
}
