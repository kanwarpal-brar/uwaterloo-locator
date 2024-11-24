import { LocationObject } from "expo-location";
import { createContext, useReducer } from "react";

export enum MapModeTypes {
  manual = "manual",
  standard = "standard",
}

export type MapContextType = {
  mode: MapModeTypes;
  lastUserLocation: LocationObject | null;
};

const defaultData: MapContextType = {
  mode: MapModeTypes.standard,
  lastUserLocation: null,
};

export const MapContext = createContext(defaultData);

export const MapDispatchContext = createContext((() => {}) as React.Dispatch<{
  type: MapActionTypes;
  payload?: any;
}>);

export function MapDataProvider({ children }: { children: React.ReactNode }) {
  const [data, dispatch] = useReducer(mapDataReducer, defaultData);
  return (
    <MapContext.Provider value={data}>
      <MapDispatchContext.Provider value={dispatch}>
        {children}
      </MapDispatchContext.Provider>
    </MapContext.Provider>
  );
}

export enum MapActionTypes {
  SET_MANUAL_MODE = "SET_MANUAL_MODE",
  SET_STANDARD_MODE = "SET_STANDARD_MODE",
  SET_USER_LOCATION = "SET_USER_LOCATION",
}

function mapDataReducer(data: MapContextType, action: any) {
  switch (action.type) {
    case MapActionTypes.SET_MANUAL_MODE:
      return { ...data, mode: MapModeTypes.manual };
    case MapActionTypes.SET_STANDARD_MODE:
      return { ...data, mode: MapModeTypes.standard };
    case MapActionTypes.SET_USER_LOCATION:
      return { ...data, lastUserLocation: action.payload };
    default:
      return data;
  }
}
