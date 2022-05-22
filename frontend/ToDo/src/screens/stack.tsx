import React from "react";
import RootStackParamList from "./rootStackParamList";
import { createStackNavigator } from "@react-navigation/stack";
import LoginScreen from "./login_screen";
import ListScreen from "./list_screen";
import ItemScreen from "./item_screen";
import UserScreen from "./user_screen";

const Stack = createStackNavigator<RootStackParamList>();

const RootStack = () => {
    return (
        <Stack.Navigator>
            <Stack.Screen name="Login" component={LoginScreen} />
            {/* <Stack.Screen name="Signup" component={Signup} /> */}
            <Stack.Screen name="List" component={ListScreen} />
            <Stack.Screen name="Item" component={ItemScreen} options={({ route }) => ({ title: route.params.list.name })} />
            <Stack.Screen name="User" component={UserScreen} options={({ title: "Update details" })} />
        </Stack.Navigator>
    );
}

export default RootStack;