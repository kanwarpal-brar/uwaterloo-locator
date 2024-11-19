import React, { useRef, useImperativeHandle, forwardRef } from "react";
import { View, StyleSheet, Text, ViewStyle } from "react-native";
import { Marker, MapMarkerProps, MapMarker, LatLng } from "react-native-maps";

export interface CustomMarkerProps extends Omit<MapMarkerProps, "ref"> {
  coordinate: LatLng;
  title?: string;
  description?: string;
  markerColor?: string;
  onPress?: () => void;
  children?: React.ReactNode;
  containerStyle?: ViewStyle;
}

export interface CustomMarkerRef {
  redraw: () => void;
  showCallout: () => void;
  hideCallout: () => void;
  animateMarkerToCoordinate: (coordinate: LatLng, duration?: number) => void;
}

const CustomMarkerV2 = forwardRef<MapMarker, CustomMarkerProps>(
  (
    {
      coordinate,
      title,
      description,
      markerColor = "#FF4444",
      onPress,
      children,
      containerStyle,
    },
    ref,
  ) => {
    const markerRef = useRef<MapMarker | null>(null);

    // useImperativeHandle(ref, () => ({
    //   redraw: () => {
    //     if (markerRef.current) {
    //       console.log("marker redraw");
    //       markerRef.current.redraw();
    //     }
    //   },
    //   showCallout: () => {
    //     if (markerRef.current) {
    //       markerRef.current.showCallout();
    //     }
    //   },
    //   hideCallout: () => {
    //     if (markerRef.current) {
    //       markerRef.current.hideCallout();
    //     }
    //   },
    //   animateMarkerToCoordinate: (coordinate: LatLng, duration?: number) => {
    //     if (markerRef.current) {
    //       markerRef.current.animateMarkerToCoordinate(
    //         coordinate,
    //         duration || 500,
    //       );
    //     }
    //   },
    // }));

    const DefaultMarkerContent: React.FC = () => (
      <View
        style={[
          styles.markerContainer,
          { borderColor: markerColor },
          containerStyle,
        ]}
      >
        <Text style={styles.markerText}>📍</Text>
      </View>
    );

    return (
      <Marker
        ref={ref}
        coordinate={coordinate}
        title={title}
        description={description}
        onPress={onPress}
        tracksViewChanges={false}
      >
        {children || <DefaultMarkerContent />}
      </Marker>
    );
  },
);

const styles = StyleSheet.create({
  markerContainer: {
    backgroundColor: "white",
    borderRadius: 8,
    borderWidth: 2,
    padding: 8,
    shadowColor: "#000",
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 3.84,
    elevation: 5,
  },
  markerText: {
    fontSize: 20,
  },
});

export default CustomMarkerV2;
