import React from "react";
import RootStackParamList from "./rootStackParamList";
import { createStackNavigator } from "@react-navigation/stack";
import LoginScreen from "./loginScreen";
import ListScreen from "./listScreen";
import ItemScreen from "./itemScreen";

const Stack = createStackNavigator<RootStackParamList>();

const RootStack = () => {
    return (
        <Stack.Navigator>
            <Stack.Screen name="Login" component={LoginScreen} />
            {/* <Stack.Screen name="Signup" component={Signup} /> */}
            <Stack.Screen name="List" component={ListScreen} />
            <Stack.Screen name="Item" component={ItemScreen} options={({ route }) => ({ title: route.params.list.name })} />
        </Stack.Navigator>
    );
}

export default RootStack;