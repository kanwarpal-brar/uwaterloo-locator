import { createContext } from "react";
import { LatLng } from "react-native-maps";

export type MapContextType = {
  settingManualLocation: boolean;
  userLocation: LatLng | null;
};

export const MapContext = createContext<MapContextType>({
  settingManualLocation: false,
  userLocation: null,
});
