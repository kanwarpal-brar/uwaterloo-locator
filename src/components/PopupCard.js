import React from "react";
import { ArrowFatLineUp, ArrowFatLineDown } from "phosphor-react";
import "../styles/Popup.css";
import { MapPin, Wheelchair, Baby, GenderIntersex } from "phosphor-react";

function PopupCard({ name, address, times, f,d,c }) {
  return (
    <div className="popup-container">
      <div className="popup-details">
        <p className="popup-details--name">{name}</p>
        <p className="popup-details--address">{address}</p>
        <p className="popup-details--times">{times}</p>

        <div style={{bottom: "0"}} className="popup-rating">
          <p style={{color: "#4188FF", fontSize: "medium", marginLeft: "0px", marginRight: "5px"}}>{parseInt((Math.random() * 100), 10)}</p>
          <ArrowFatLineUp size={27} color="#4188FF" weight="fill" />
          <p style={{ color: "#EE6462", fontSize: "medium", marginLeft: "10px", marginRight: "5px"}}>{parseInt((Math.random() * 100), 10)}</p>
          <ArrowFatLineDown size={27} color="#EE6462" weight="fill" />
        </div>
        

      </div>
      
      <div className="popup-rating">
      {/* <div class="vl"></div> */}
      <div className='icon-container'> 
                    {f==1? <div className='outside'><GenderIntersex size={32} color="#ffffff" weight="fill" /></div> : <></>}
                    {d==0? <></> : <div className='outside'><Wheelchair size={32} color="#ffffff" weight="fill" /></div>}
                    {c==1? <div className='outside'><Baby size={32} color="#ffffff" weight="fill" /></div> : <></>}
        </div>
        
      </div>
    </div>
  );
}

<div flex justify space between>
  <div>
    <p></p>
    <p></p>
  </div>
  <div></div>
</div>;

export default PopupCard;
