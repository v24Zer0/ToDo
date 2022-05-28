import List from "../models/list";
import User from "../models/user";

type RootStackParamList = {
    Login: undefined;
    List: undefined;
    Item: {list: List}; 
    User: {user: User};
    Signup: undefined;
}

export default RootStackParamList;