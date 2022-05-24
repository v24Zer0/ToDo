import React from "react";
import { Pressable, Text, View } from "react-native";
import Item from "../models/item";
import List from "../models/list";

interface ItemProps {
    item: Item;
    setModalVisible: React.Dispatch<React.SetStateAction<boolean>>;
    setModalItem: React.Dispatch<React.SetStateAction<Item>>;
}

// Add Pressable component to trigger modal display
// onPress triggers Modal
const ItemComponent: React.FC<ItemProps> = ({ item, setModalVisible, setModalItem }) => {
    return (
        <View>
            <Pressable onPress={() => {
                setModalVisible(true);
                setModalItem(item);
            }}>
                <Text>
                    {item.task}
                </Text>
            </Pressable>
        </View>
    );
}

export default ItemComponent;