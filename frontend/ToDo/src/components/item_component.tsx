import React from "react";
import { Text, View } from "react-native";
import Item from "../models/item";

const ItemComponent: React.FC<Item> = ({ id, task, priority, list_id }) => {
    return (
        <View>
            <Text>
                {task}
            </Text>
        </View>
    );
}

export default ItemComponent;