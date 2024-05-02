import { Children, ReactNode, cloneElement, isValidElement } from "react";
import { View } from "react-native";
import Svg, { G, Circle, Path } from "react-native-svg";

const aspectWidth = 20;
const aspectHeight = 34.892337;
const ratio = aspectWidth / aspectHeight;

export default function CustomPin({
  icon,
  fill,
  stroke,
  strokeWidth,
}: {
  icon?: any;
  fill?: string;
  stroke?: string;
  strokeWidth?: number;
}) {
  return (
    <View style={{ aspectRatio: ratio }}>
      <Svg height={45} viewBox={`0 0 ${aspectWidth} ${aspectHeight}`}>
        <G transform="translate(-814.59595,-274.38623)">
          <G transform="matrix(1.1855854,0,0,1.1855854,-151.17715,-57.3976)">
            <Path
              d="m 817.11249,282.97118 c -1.25816,1.34277 -2.04623,3.29881 -2.01563,5.13867 0.0639,3.84476 1.79693,5.3002 4.56836,10.59179 0.99832,2.32851 2.04027,4.79237 3.03125,8.87305 0.13772,0.60193 0.27203,1.16104 0.33416,1.20948 0.0621,0.0485 0.19644,-0.51262 0.33416,-1.11455 0.99098,-4.08068 2.03293,-6.54258 3.03125,-8.87109 2.77143,-5.29159 4.50444,-6.74704 4.56836,-10.5918 0.0306,-1.83986 -0.75942,-3.79785 -2.01758,-5.14062 -1.43724,-1.53389 -3.60504,-2.66908 -5.91619,-2.71655 -2.31115,-0.0475 -4.4809,1.08773 -5.91814,2.62162 z"
              fill={`${fill || "#ff4646"}`}
              stroke={`${stroke || "#d73534"}`}
              strokeWidth={strokeWidth || 0}
              strokeMiterlimit={4}
              strokeDasharray={"none"}
            />
            {icon ? (
              <View
                style={{
                  width: "75%",
                  aspectRatio: 1,
                  alignSelf: "center",
                  top: "33%",
                }}
              >
                {icon}
              </View>
            ) : (
              <Circle
                r="3.0355"
                cy={"288.25278"}
                cx={"823.03064"}
                fill={"#560000"}
                strokeWidth={0}
              />
            )}
          </G>
        </G>
      </Svg>
    </View>
  );
}
