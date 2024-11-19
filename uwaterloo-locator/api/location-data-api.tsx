import buildings from "../constants/buildings.json";

export type NamedLatLng = {
  name: string;
  id?: string;
  latitude: number;
  longitude: number;
};

export function fetchWashroomLocations(): Array<NamedLatLng> {
  // TODO this is currently mocked by parsing a geojson: connect to backend later
  return buildings.features.map((build) => {
    return {
      name: build.properties.buildingName,
      buildingId: build.properties.buildingId,
      latitude: build.properties.latitude || 0,
      longitude: build.properties.longitude || 0,
    };
  });
}
