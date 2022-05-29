import React from "react";
import { Pressable, Text, View } from "react-native";
import Item from "../models/item";

interface ItemProps {
    item: Item;
    setModalVisible: React.Dispatch<React.SetStateAction<boolean>>;
    setModalItem: React.Dispatch<React.SetStateAction<Item>>;
}

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