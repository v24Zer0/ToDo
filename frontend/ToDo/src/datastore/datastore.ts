/* 
    Datastore for login and auth data implemented using react-native-async-storage 
*/

import AsyncStorage from "@react-native-async-storage/async-storage";

const storeAuthToken = async (token: string) => {
    try {
        await AsyncStorage.setItem("token", token);
    } catch(e) {
        console.log("Error storing token", e);
    }
}

const getAuthToken = async () => {
    try {
        const token = await AsyncStorage.getItem("token");
        if(token !== null) {
            return token;
        }
        return "";
    } catch(e) {
        return "";
    }
}

const storeUserInfo = (userID: string, username: string, loggedIn: boolean) => {

}

const getUserInfo = () => {
    return "";
}

export default {
    storeAuthToken,
    getAuthToken,
    storeUserInfo,
    getUserInfo
};