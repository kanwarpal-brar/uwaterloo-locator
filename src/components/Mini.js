
import '../styles/Mini.css'
import { MapPin, Toilet } from "phosphor-react";
export function Card(){
    return (
        <div className="mini-info">
            <p className='loo'>UWaterLoo Locator</p>
            <Toilet size={32} color="#000000" weight="fill" />
        </div>
    )
}

export default Card;