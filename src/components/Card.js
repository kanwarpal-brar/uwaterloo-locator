
import '../styles/Card.css'
import { MapPin, Wheelchair, Baby, GenderIntersex } from "phosphor-react";

function haversineDistance(coords1, coords2, isMiles) {
    function toRad(x) {
      return x * Math.PI / 180;
    }
  
    var lon1 = coords1[0];
    var lat1 = coords1[1];
  
    var lon2 = coords2[0];
    var lat2 = coords2[1];
  
    var R = 6371; // km
  
    var x1 = lat2 - lat1;
    var dLat = toRad(x1);
    var x2 = lon2 - lon1;
    var dLon = toRad(x2)
    var a = Math.sin(dLat / 2) * Math.sin(dLat / 2) +
      Math.cos(toRad(lat1)) * Math.cos(toRad(lat2)) *
      Math.sin(dLon / 2) * Math.sin(dLon / 2);
    var c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));
    var d = R * c;
  
    if(isMiles) d /= 1.60934;

    return d;
  }


export function Card(washroom) {
    console.log(washroom);
    return (
        <div className="info-card">
            {/* <img src="/images/Washroom.jpeg" alt="" className='image'></img> */}
            <div className="info-area">
            <div className="top-area">
                <MapPin size={16} weight="fill" />
                <p className="distance">{Math.floor(haversineDistance([washroom.washroom.properties.X_COORDINATE,washroom.washroom.properties.Y_COORDINATE],[-75.711210,45.398930],false)) + "km"}</p>
            </div>
            <h1 className='name'>{washroom.washroom.properties.NAME}</h1>
            <div class="line"></div>
            <p className='address'>{washroom.washroom.properties.ADDRESS}</p>
            <p className='time'>{washroom.washroom.properties.HOURS_MONDAY_OPEN + " - " + washroom.washroom.properties.HOURS_MONDAY_CLOSED }</p>
            {/* <p> 0.1km </p> */}
            <div className='icon-container'> 
                    {washroom.washroom.properties.FAMILY_TOILET==1? <div className='outside'><GenderIntersex size={32} color="#ffffff" weight="fill" /></div> : <></>}
                    {washroom.washroom.properties.SPECIAL_TOILET_TYPE==0? <></> : <div className='outside'><Wheelchair size={32} color="#ffffff" weight="fill" /></div>}
                    {washroom.washroom.properties.CHANGE_STATION_CHILD==1? <div className='outside'><Baby size={32} color="#ffffff" weight="fill" /></div> : <></>}
            </div>

            </div>
        </div>
    )
}

export default Card;