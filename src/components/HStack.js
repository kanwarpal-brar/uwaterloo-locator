
import '../styles/HStack.css'
import Card from './Card';
import washroomData from '../data/Public_Washrooms.json'
export function HStack(){
    return (
        <div className="HStack">
        {washroomData.features.map((washroom,index) => {
            var latitude = washroom.properties.Y_COORDINATE;
            var longitude = washroom.properties.X_COORDINATE;
            // console.log(washroom.properties.Y_COORDINATE, washroom.properties.X_COORDINATE)
            return (
                <Card washroom={washroom}/>
            );
          })}
        </div>
    )
}

export default HStack;