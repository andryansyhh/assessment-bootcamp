import {REGISTER_SUCCESS, REGISTER_FAIL} from "../action/actionType"
import axios from 'axios'
import { API_URL } from "../../utils/utils"

export const regiterAPI = () => async (dispatch) => {
    await axios.post(`${API_URL}/users/register`)
    
    dispatch({
        type:REGISTER_SUCCESS
    })

    .catch((err)=>{
        console.log(err.message)
        dispatch({
            type:REGISTER_FAIL
        })
    });

}

