import { View } from "react-native";
import Svg, { G, Circle, Path } from "react-native-svg";
import { FontAwesome6 } from "@expo/vector-icons";

const aspectWidth = 20;
const aspectHeight = 34.892337;
const ratio = aspectWidth / aspectHeight;

export type CustomIconType =
  | "toilet"
  | "person"
  | "person-dress"
  | "person-half-dress";

const iconMap: Record<CustomIconType, React.ReactNode> = {
  toilet: (
    <FontAwesome6
      name="toilet"
      size={15}
      color="#000000"
      style={{
        alignSelf: "center",
        transform: [{ translateX: 0.31 }, { translateY: 0 }],
      }}
    />
  ),
  person: (
    <FontAwesome6
      name="person"
      size={20}
      color="#ededed"
      style={{
        alignSelf: "center",
        transform: [{ translateX: 0.31 }, { translateY: -1.5 }],
      }}
    />
  ),
  "person-dress": (
    <FontAwesome6
      name="person-dress"
      size={20}
      color="#ededed"
      style={{
        alignSelf: "center",
        transform: [{ translateX: 0.31 }, { translateY: -1.5 }],
      }}
    />
  ),
  "person-half-dress": (
    <FontAwesome6
      name="person-half-dress"
      size={20}
      color="#ededed"
      style={{
        alignSelf: "center",
        transform: [{ translateX: 0.31 }, { translateY: -1.5 }],
      }}
    />
  ),
};

export type CustomPinProps = {
  icon?: CustomIconType;
  fill?: string;
  stroke?: string;
  strokeWidth?: number;
};

export default function CustomPin({
  icon,
  fill,
  stroke,
  strokeWidth,
}: CustomPinProps) {
  return <View>{iconMap["toilet"]}</View>;
}
