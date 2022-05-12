import { createNativeStackNavigator } from "@react-navigation/native-stack";
import ItemList from "./components/item_list";
import Login from "./components/login";
import Signup from "./components/signup";
import UserLists from "./components/user_lists";

const Stack = createNativeStackNavigator();

const MyStack = () => {
    return (
        <Stack.Navigator>
            <Stack.Screen name="Login" component={Login} />
            <Stack.Screen name="Signup" component={Signup} />
            <Stack.Screen name="Lists" component={UserLists} />
            <Stack.Screen name="Items" component={ItemList} />
        </Stack.Navigator>
    );
}

export default MyStack;