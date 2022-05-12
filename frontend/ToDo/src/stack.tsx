import { createNativeStackNavigator } from "@react-navigation/native-stack";
import ItemList from "./components/item_list";
import UserLists from "./components/user_lists";

const Stack = createNativeStackNavigator();

const MyStack = () => {
    return (
        <Stack.Navigator>
            <Stack.Screen name="Lists" component={UserLists} />
            <Stack.Screen name="Items" component={ItemList} />
        </Stack.Navigator>
    );
}

export default MyStack;