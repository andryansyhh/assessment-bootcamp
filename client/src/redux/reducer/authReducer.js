import {REGISTER_SUCCESS, REGISTER_FAIL} from '../action/actionType'

const user =JSON.parse(localStorage.getItem("user"))

const initialState = 
user ? {
    isLogin :true, user
}:
{
    isLogin:false
}

const authReducer = (state = initialState, action) => {
    if (action.type === REGISTER_SUCCESS){
        return{
            ...state,
            isLogin:false
        }
    }

    if (action.type === REGISTER_FAIL){
        return{
            ...state,
            isLogin:false
        }
    }
}

export default authReducer